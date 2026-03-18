import { matchesSentencePattern } from './sentenceIndex';

describe('matchesSentencePattern', () => {
  it('matches a step with a single {string} wildcard', () => {
    expect(
      matchesSentencePattern(
        'the response should have field {string}',
        'the response should have field "user.name"',
      ),
    ).toBe(true);
  });

  it('matches a step with two {string} wildcards', () => {
    expect(
      matchesSentencePattern(
        'the user enters {string} into the {string} field',
        'the user enters "admin" into the "username" field',
      ),
    ).toBe(true);
  });

  it('does not match when literal text differs', () => {
    expect(
      matchesSentencePattern(
        'the response should have field {string}',
        'the request should have field "user.name"',
      ),
    ).toBe(false);
  });

  it('matches an {int} wildcard against a plain number', () => {
    expect(
      matchesSentencePattern(
        'the response status code should be {int}',
        'the response status code should be 200',
      ),
    ).toBe(true);
  });

  it('is case-insensitive', () => {
    expect(
      matchesSentencePattern(
        'I send the request',
        'I SEND THE REQUEST',
      ),
    ).toBe(true);
  });
});
