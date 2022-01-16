export async function handler(event, context) {
  console.log(`Function name: ${context.functionName}`)
  console.log(`Event: ${JSON.stringify(event)}`)

  const response = {
    statusCode: 200,
    "headers": {
      "Content-Type": "application/json"
    },
    "body": JSON.stringify("Hello, world!")
  };
  return response;
};
