// Minimal TypeScript stub for the taikuri testbed front-end.
//
// Future map-edge tests assert that frontend/ must NOT import a backend-internal
// package (e.g. ../internal/...). No such import exists yet — this file is the
// clean baseline those tests start from.
export function greet(name: string): string {
  return `hello, ${name}`;
}

console.log(greet("taikuri"));
