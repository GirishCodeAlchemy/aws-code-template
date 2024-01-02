// Lambda handler function
exports.handler = async (event) => {
    // Print the received event
    console.log('Received Event:', JSON.stringify(event));


    const message = `Recieved message, ${event}`;

    // Creating a response
    const response = {
        statusCode: 200,
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ message }),
    };

    return response;
};
