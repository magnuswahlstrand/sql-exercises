import {CfnOutput, DockerImage, Stack, StackProps} from 'aws-cdk-lib';
import * as go_alpha from '@aws-cdk/aws-lambda-go-alpha'


import {Construct} from 'constructs';
import {Architecture, FunctionUrlAuthType} from "aws-cdk-lib/aws-lambda";

export class InfraStack extends Stack {
    constructor(scope: Construct, id: string, props?: StackProps) {
        super(scope, id, props);


        const fn = new go_alpha.GoFunction(this, 'Main', {
            architecture: Architecture.ARM_64,
            entry: '../functions/lambda-entrypoint',
            bundling: {
                cgoEnabled: true,
                forcedDockerBundling: true,
                dockerImage: DockerImage.fromBuild("lib"),
                environment: {
                    "GOMODCACHE": "/tmp/",
                    "GOCACHE": "/tmp/",
                    "GOFLAGS": "-buildvcs=false"
                }
            }
        })

        const fnUrl = fn.addFunctionUrl({
            authType: FunctionUrlAuthType.NONE
        })


        // Output the function URL
        new CfnOutput(this, 'FunctionUrlOutput', {
            value: fnUrl.url,
        });
    }
}
