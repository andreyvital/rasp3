/* @flow */
import type { MP3 } from './MP3'
import React, { Component } from 'react'
import AlbumList from './AlbumList'
import fetchQuery from './fetchQuery'
import './linea.css'
import './rasp3.css'

export default class Root extends Component {
  state: {
    mp3: ?MP3[]
  };

  constructor(props: any) {
    super(props)
    this.state = {
      mp3: []
    }
  }

  componentDidMount(): void {
    fetchQuery(
      `
{
  mp3 {
    file
    size
    id3 {
      title
      artist
      track
      album
      year
    }
  }
}
      `
    ).then(({ mp3 }) => {
      this.setState({
        mp3: mp3.slice(0, 30)
      })
    })
  }

  render(): ?ReactElement {
    return (
      <AlbumList mp3={this.state.mp3} />
    )
  }
}
