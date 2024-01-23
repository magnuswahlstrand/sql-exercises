import {Api, Config, StackContext, Table} from "sst/constructs";

export function ApiStack({stack}: StackContext) {
    const api = new Api(stack, "SqlExercises", {
        defaults: {
            function: {
                environment: {}
            }
        },
        routes: {
            "ANY /{proxy+}": "functions/lambda-entrypoint/main.go",
            "OPTIONS /{proxy+}": "functions/lambda-entrypoint/main.go",
        },

    });
    stack.addOutputs({
        "ApiEndpoint": api.url,
    });
}
