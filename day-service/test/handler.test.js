// /test/handler/test.js

const LambdaTester = require('lambda-tester');
const { calculateDay } = require( '../handler' );

const testEvent = {
  body: '{"date":"22 May 2017"}',
};

describe('handler()', () => {
  test('returns the correct day of the week', () =>
    LambdaTester(calculateDay) // Which lambda function to test
      .event(testEvent) // Supply a test event
      .expectResult((response) => { // Test assertions
        const data = JSON.parse(response.body);
        expect(data.day).toBe('Monday');
      })
  );
});
