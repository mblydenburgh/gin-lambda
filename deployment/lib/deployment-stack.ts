import cdk = require('@aws-cdk/core');
import { Code, Function, Runtime } from "@aws-cdk/aws-lambda"
import { Effect, ManagedPolicy, PolicyStatement, Role, ServicePrincipal } from '@aws-cdk/aws-iam';
import { CorsHttpMethod, HttpApi, HttpMethod } from "@aws-cdk/aws-apigatewayv2"
import { Table, AttributeType, BillingMode, ProjectionType } from "@aws-cdk/aws-dynamodb"
import { HttpLambdaIntegration } from "@aws-cdk/aws-apigatewayv2-integrations"

export class DeploymentStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const appName = "gin-lambda"

    //@ts-ignore
    const dynamoTable = new Table(this, `DynamoTable`, {
            tableName: `${appName}-table`,
            partitionKey: { name: "UserId", type: AttributeType.STRING },
            sortKey: { name: "ModelTypeAndId", type: AttributeType.STRING },
            billingMode: BillingMode.PAY_PER_REQUEST
    })
//     dynamoTable.addGlobalSecondaryIndex({
//      indexName: "vinIndex",
//      projectionType: ProjectionType.ALL,
//     partitionKey: {name: "VIN", type: AttributeType.STRING},
//    })



    //@ts-ignore
    const lambdaRole = new Role(this, `LambdaRole`, {
            roleName: `${appName}-role`,
            assumedBy: new ServicePrincipal("lambda.amazonaws.com"),
            managedPolicies: [
              ManagedPolicy.fromAwsManagedPolicyName("service-role/AWSLambdaVPCAccessExecutionRole")
            ]
    })
    dynamoTable.grantReadWriteData(lambdaRole)
        lambdaRole.addToPolicy(new PolicyStatement({
            effect: Effect.ALLOW,
            actions: [
                "dynamodb:Query",
                "dynamodb:Scan",
                "dynamodb:GetItem",
                "dynamodb:PutItem",
                "dynamodb:UpdateItem",
                "dynamodb:DeleteItem"
            ],
            resources: [
                `arn:aws:dynamodb:${cdk.Aws.REGION}:${cdk.Aws.ACCOUNT_ID}:table/${dynamoTable.tableName}/index/*`
            ]
        }))

    //@ts-ignore
    const lambdaFunction = new Function(this, `LambdaFunction`, {
            functionName: `${appName}-lambda`,
            runtime: Runtime.GO_1_X,
            role: lambdaRole,
            code: Code.fromAsset("../src/bin/main.zip"),
            handler: "main",
            environment: {
              "TABLE_NAME": "gin-lambda-table"
            }
    })

    lambdaFunction.grantInvoke(new ServicePrincipal("apigateway.amazonaws.com"))

    const lambdaIntegration = new HttpLambdaIntegration("HttpLambdaIntegration", lambdaFunction)

    //@ts-ignore
    const api = new HttpApi(this, `RestAPIGateway`, {
            apiName: "gin-api",
            corsPreflight: {
                allowHeaders: ['Authorization', 'Access-Control-Allow-Origin', 'Access-Control-Allow-Headers', 'Content-Type', "X-Api-Key", "X-Amz-Security-Token"],
                allowMethods: [
                    CorsHttpMethod.ANY
                ],
                allowOrigins: ['*'],
            },
        })

        api.addRoutes({
            path: "/",
            methods: [HttpMethod.ANY],
            integration: lambdaIntegration
        })

        api.addRoutes({
            path: "/{proxy+}",
            methods: [HttpMethod.ANY],
            integration: lambdaIntegration
        })

  }
}
