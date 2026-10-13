'use strict'

const { buildProgram, starterCode } = require('./harness')

const sampleProblem = {
  function: {
    name: 'twoSum',
    pythonName: 'two_sum',
    params: [
      { name: 'nums', type: 'integer[]' },
      { name: 'target', type: 'integer' }
    ],
    returns: 'integer[]'
  }
}

test('generates valid Go starter code', () => {
  const code = starterCode(sampleProblem, 'go')
  expect(code).toContain('func twoSum(nums []int, target int) []int')
})

test('generates program with package main at the top', () => {
  const program = buildProgram({
    language: 'go',
    code: 'func twoSum(nums []int, target int) []int { return []int{0, 1} }',
    problem: sampleProblem,
    nonce: '0123456789abcdef0123456789abcdef'
  })
  expect(program.trim().startsWith('package main')).toBe(true)
  expect(program).toContain('json.RawMessage')
  expect(program).toContain('defer func() {')
  expect(program).toContain('recover()')
})