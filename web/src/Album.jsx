/* @flow */
import type { MP3 } from './MP3'
import { sortBy } from 'lodash'
import React, { Component } from 'react'
import TrackRow from './TrackRow'

let AlbumStyles = {
  title: {
    padding: 10,
    paddingLeft: 0,
    margin: 0,
    color: '#FFF',
    fontWeight: 'lighter',
    fontSize: 28
  }
}

export default class Album extends Component {
  props: {
    album: string;
    mp3: MP3[];
  };

  _guessAlbumYear(mp3: MP3[]): ?number {
    if (mp3[0].id3) {
      return ((mp3[0].id3: $FlowIssue).year: $FlowIssue)
    }

    return null
  }

  render(): Component {
    let year = this._guessAlbumYear(this.props.mp3)

    return (
      <div>
        <h3 style={AlbumStyles.title}>
          {this.props.album}
          {year && (
            <span>{' '}&mdash;{' '}{year}</span>
          )}
        </h3>
        {sortBy(this.props.mp3, mp3 => mp3.id3.track).map(mp3 => {
          return (
            <TrackRow
              key={mp3.file}
              mp3={mp3} />
          )
        })}
      </div>
    )
  }
}
