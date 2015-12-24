/* @flow */
import React, { Component } from 'react'

export default class PlayTrackNumber extends Component {
  props: {
    track: number;
    onPlay: (track: number) => void;
  };

  state: {
    hover: boolean
  };

  constructor(props: any) {
    super(props)

    this.state = {
      hover: false
    }
  }

  mouseEnter(e: SyntheticMouseEvent): void {
    this.setState({
      hover: true
    })
  }

  mouseOut(e: SyntheticMouseEvent): void {
    this.setState({
      hover: false
    })
  }

  render(): Component {
    return (
      <div
        onMouseEnter={this.mouseEnter.bind(this)}
        onMouseLeave={this.mouseOut.bind(this)}
        style={this.state.hover ? {
          padding: 8,
          textAlign: 'center'
        } : {
          border: '1px solid #EB5E8A',
          padding: 8,
          borderRadius: 16
        }}>
        <strong
          style={{
            fontSize: 18,
            color: '#FFF',
            display: this.state.hover ? 'none' : null
          }}>
          {this.props.track < 10 ? '0' + this.props.track : this.props.track}
        </strong>
        {this.state.hover && (
          <div
            onClick={this.props.onPlay.bind(null, this.props.track)}
            style={{
              fontSize: 23,
              color: '#FFF',
              cursor: 'pointer',
              padding: 0,
              margin: 0
            }}
            className="icon icon-music-play-button">
          </div>
        )}
      </div>
    )
  }
}
