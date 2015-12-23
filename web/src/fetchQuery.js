/* @flow */
import fetch from 'isomorphic-fetch'

export default function fetchQuery(query: string): Promise<any> {
  return new Promise((resolve, reject) => {
    fetch(`http://localhost:2015/query`, {
      method: 'POST',
      body: query,
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/graphql'
      }
    }).catch(reject)
      .then(response => response.json())
      .then(response => {
        resolve(response)
      })
  })
}
