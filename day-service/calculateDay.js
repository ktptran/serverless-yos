'use strict';

const getDayFromDate = (dateString) => {
  const DAYS = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday',\
  'Saturday'];
  const date = new Date(Date.parse(dateString
  if (date === 'Invalid Date') {
    throw new Error('Date parsing has failed!');
  }
  const dayIndex = date.getDay();
  const day = DAYS[dayIndex];
  return day;
}

// Handler Function
module.exports.calculateDay = (event, context, callback) { // Renamed to calculateDay
  console.log('calculateDelay Invoked!')
  const input = JSON.parse(event.body); // 1. Get http body from the elemtn
  console.log(event); // Log event
  console.log(input); // Log HTTP body
  let day;
  try {
    day = getDayFromDate(input.date);
  } catch (e) {
    const errorResponse = {
      statusCode: 500,
      headers: {
        'Access-Control-Allow-Origin': '*', // Required for CORS support to work
      },
      body: JSON.stringify({
        error: e.message, // Specific error message
      }),
    };
    callback(null, errorResponse);
    return;
  }
  const successResponse = {
    statusCode: 200,
    headers: {
      'Access-Control-Allow-Origin': '*', // Required for CORS support to work
    },
    body: JSON.stringify({
      date: input.date,
      day: getDayFromDate(input.date), // 2. Return result from getDayFromDate
    }),
  };
  callback(null, response);
}
