/* @flow */
import React, { Component } from 'react'
import PlayTrackNumber from './PlayTrackNumber'

let TrackRowStyles = {
  container: {
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    width: 300,
    marginTop: 2,
    marginBottom: 2
  }
}

export default class TrackRow extends Component {
  render(): ?ReactElement {
    let { id3: { title, artist, album, track } } = this.props.mp3

    return (
      <div style={TrackRowStyles.container}>
        <PlayTrackNumber onPlay={() => {}} track={track} />
        <div style={{ padding: 10 }}>
          <div
            style={{
              fontSize: 15,
              fontWeight: 'bold',
              color: '#FFF'
            }}>
            {title}
          </div>
          <div style={{ fontSize: 14 }}>
            <em>{album}</em> &mdash; {artist}
          </div>
        </div>
      </div>
    )
  }
}
