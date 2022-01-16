import minimist from "minimist";

export const cliArgs = minimist(process.argv)
export const baseUrl = cliArgs.baseUrl ?? "http://localhost:3000"
