#!/usr/bin/env node
import 'source-map-support/register';
import cdk = require('@aws-cdk/core'); 
import { DeploymentStack } from '../lib/deployment-stack';

const app = new cdk.App();
new DeploymentStack(app, 'DeploymentStack', {
   env: { account: '415023725722', region: 'us-east-1' },
});
