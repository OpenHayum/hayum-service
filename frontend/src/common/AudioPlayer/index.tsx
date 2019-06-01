import React, {Component, ReactElement} from "react";
import classnames from "classnames";
import * as PropTypes from "prop-types";

import "./audioPlayer.less";


enum controlNames {
  IsPlaying = 'isPlaying'
}

interface AudioPlayerProps {
  // TODO: make src `required` when AudioPlayer is more stable
  // it is just for testing purpose
  src?: string;
}

interface AudioPlayerState {
  playheadPosition: number;
  activeTimelineWidth: number;
  currentDuration: string;
  totalDuration: string;
  [controlNames.IsPlaying]?: boolean;
}

class AudioPlayer extends Component<AudioPlayerProps, AudioPlayerState> {
  static propTypes = {
    src: PropTypes.string
  };

  static defaultProps = {
    src: "http://www.panacherock.com/downloads/mp3/01_Afterlife.mp3"
  };

  state: AudioPlayerState;
  timeline: HTMLElement | null;
  playhead: HTMLElement | null;
  player: HTMLAudioElement | null;
  mouseOnPlayhead: boolean;
  timelineWidth: string | null;
  duration: any;
  handleEnded: any;

  constructor(props) {
    super(props);
    this.state = {
      isPlaying: false,
      playheadPosition: 0,
      activeTimelineWidth: 0,
      currentDuration: "00:00",
      totalDuration: this.formatTime(0)
    };

    this.timeline = null;
    this.playhead = null;
    this.timelineWidth = null;
    this.mouseOnPlayhead = false;
    this.duration = null;
    this.player = null;
  }

  componentDidMount(): void {
    window.addEventListener("mouseup", this.mouseUp, false);
  }

  componentWillUnmount(): void {
    window.removeEventListener("mouseup", this.mouseUp, false);
    if (!this.player) return;

    this.player.removeEventListener('canplaythrough', this.handleCanPlayThrough, false);
    this.player.removeEventListener('timeupdate', this.handleTimeUpdate, false);
    this.player.removeEventListener('ended', this.handleEnded, false);
  }

  clickPercent = (event: MouseEvent): number => {
    if (!this.timeline) return 0;
    const {left, width} = this.getPosition(this.timeline);
    return (event.clientX - left) / width;
  };

  formatTime = (seconds: number): string => {
    if (!seconds) return "00:00";
    const minutes = Math.floor(seconds / 60);
    const minutesStr = minutes >= 10 ? minutes + "" : "0" + minutes;
    seconds = Math.floor(seconds % 60);
    const secondsStr = seconds >= 10 ? seconds + "" : "0" + seconds;
    return minutesStr + ":" + secondsStr;
  };

  getPosition = (el: HTMLElement): ClientRect => {
    return el.getBoundingClientRect();
  };

  handleControlClick = (): void => {
    if (!this.player) return;

    const {isPlaying} = this.state;

    if (isPlaying) {
      this.player.pause();
    } else {
      this.player.play();
    }

    this.setState({
      ...this.state,
      isPlaying: !isPlaying
    });
  };

  handleTimelineClick = (e: MouseEvent): void => {
    // console.info(e);
    this.movePlayhead(e);
  };

  handleTimeUpdate = (): void => {
    if (!this.player) return;
    const timelineWidth = 300;
    const playPercent = timelineWidth * (this.player.currentTime / this.duration);

    this.setState({
      playheadPosition: playPercent,
      currentDuration: this.formatTime(this.player.currentTime),
    });
  }

  handleCanPlayThrough = (): void => {
    if (!this.player) return;
    this.setState({totalDuration: this.formatTime(this.player.duration)})
  }

  mouseDown = (): void => {
    this.mouseOnPlayhead = true;
    window.addEventListener("mousemove", this.movePlayhead, true);
  };

  mouseUp = (event: MouseEvent): void => {
    if (this.mouseOnPlayhead) {
      this.movePlayhead(event);
      window.removeEventListener("mousemove", this.movePlayhead, true);
    }
    this.mouseOnPlayhead = false;
  };

  movePlayhead = (event: MouseEvent): void => {
    if (!this.timeline) return;

    const {left, width: timelineWidth} = this.getPosition(this.timeline);
    // const { duration } = this.state;
    var newMargLeft = event.clientX - left;
    let playheadPosition = newMargLeft;

    if (newMargLeft < 0) {
      playheadPosition = 0;
    }
    if (newMargLeft > timelineWidth) {
      playheadPosition = timelineWidth;
    }

    this.setState({
      playheadPosition,
      activeTimelineWidth: playheadPosition,
      currentDuration: this.formatTime(
          this.duration * (playheadPosition / timelineWidth)
      )
    });
  };

  render(): ReactElement {
    const {
      isPlaying,
      playheadPosition,
      activeTimelineWidth,
      currentDuration,
      totalDuration
    } = this.state;

    return (
        <div className="ap">
          <div className="ap__item">
            <div className="ap__item__details">
              <h4 className="ap__item__details__name">Loktak Ema</h4>
              <h4 className="ap__item__details__artist">Ranbir</h4>
            </div>
          </div>
          <div className="ap__container">
            <div className="ap__container__wrapper">
              <div className="ap__container__controls">
                <button className="ap__icon ap__control ap__right-spacing">
                  <i className="icon-shuffle"/>
                </button>
                <button className="ap__icon ap__control ap__right-spacing">
                  <i className="icon-control-start"/>
                </button>
                <button
                    name={controlNames.IsPlaying}
                    className={classnames(
                        "ap__icon ap__circle ap__control",
                        {
                          "ap__circle--play": isPlaying,
                          "ap__circle--pause": !isPlaying
                        }
                    )}
                    onClick={this.handleControlClick}
                />
                <button className="ap__icon ap__control ap__left-spacing">
                  <i className="icon-control-end"/>
                </button>
                <button className="ap__icon ap__control ap__left-spacing">
                  <i className="icon-loop"/>
                </button>
              </div>
              <div className="ap__timeline-container">
                <div
                    className="ap__playback-time ap__playback-time__left ap__control">
                  {currentDuration}
                </div>
                <div className="ap__timeline" ref={this.setTimelineRef}>
                  <div
                      className="ap__timeline__current"
                      style={{width: activeTimelineWidth}}
                  />
                  <div
                      className="ap__playhead"
                      ref={this.setPlayheadRef}
                      style={{left: playheadPosition}}
                  />
                </div>
                <div className="ap__playback-time ap__playback-time__right ap__control">
                  {totalDuration}
                </div>
              </div>
              <audio ref={this.setPlayerRef} src={this.props.src}/>
            </div>
          </div>
          <div className="ap__volume">
            <div>
              <div className="ap__volume__icon">
                <i className="icon-volume-2"/>
              </div>
              <div className="ap__volume__controls">
                <div className="ap__volume__controls__timeline"/>
              </div>
            </div>
          </div>
        </div>
    );
  }

  setTimelineRef = _ref => {
    this.timeline = _ref;
    if (!this.timeline) return;
    this.timelineWidth = this.timeline.style.width;
    this.timeline.addEventListener("click", this.handleTimelineClick, false);
  };

  setPlayheadRef = _ref => {
    this.playhead = _ref;
    if (!this.playhead) return;
    this.playhead.addEventListener("mousedown", this.mouseDown, false);
  };

  setPlayerRef = _ref => {
    this.player = _ref;
    if (!this.player) return;
    this.player.volume = 0.2;
    this.player.addEventListener('canplaythrough', this.handleCanPlayThrough, false);
    this.player.addEventListener('timeupdate', this.handleTimeUpdate, false);
    this.player.addEventListener('ended', this.handleEnded, false);
  };
}

export default AudioPlayer;