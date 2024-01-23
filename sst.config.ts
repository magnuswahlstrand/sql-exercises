import {SSTConfig} from "sst";
import { ApiStack } from "./stacks/ApiStack";

export default {
    config(_input) {
        return {
            name: "sql-exercises",
            region: "eu-north-1",

        };
    },
    stacks(app) {
        app.setDefaultFunctionProps({
            runtime: "go",
        });
        app.stack(ApiStack);
    }
} satisfies SSTConfig;
