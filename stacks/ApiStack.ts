import {Api, StackContext, Function} from "sst/constructs";

export function ApiStack({stack}: StackContext) {
    const routerFunc = new Function(stack, "Router", {
        handler: "functions/lambda-entrypoint/main.go",
        go: {
            // cgoEnabled: true,
            // ldFlags: ["-linkmode external -extldflags \"-static\""],

        }
    });

    const api = new Api(stack, "SqlExercises", {
        routes: {
            "ANY /{proxy+}": routerFunc,
            "OPTIONS /{proxy+}": routerFunc,
        },

    });
    stack.addOutputs({
        "ApiEndpoint": api.url,
    });
}
