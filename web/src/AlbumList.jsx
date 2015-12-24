/* @flow */
import type { MP3 } from './MP3'
import { groupBy } from 'lodash'
import React, { Component } from 'react'
import Album from './Album'

let AlbumListStyles = {
  separator: {
    border: '0px',
    borderTop: '1px solid rgba(0, 0, 0, 0.1)'
  }
}

export default class AlbumList extends Component {
  props: {
    mp3: MP3[];
  };

  render(): ?ReactElement {
    let albums = groupBy(this.props.mp3.filter(mp3 => {
      return mp3.id3 && mp3.id3.album
    }), mp3 => {
      return mp3.id3.album.toLowerCase()
    })

    let names = Object.keys(albums)

    return (
      <div>
        {names.map((album, index) => {
          return (
            <div key={album}>
              <Album
                album={albums[album][0].id3.album}
                mp3={albums[album]} />
              {index < names.length && (
                <hr style={AlbumListStyles.separator} />
              )}
            </div>
          )
        })}
      </div>
    )
  }
}
