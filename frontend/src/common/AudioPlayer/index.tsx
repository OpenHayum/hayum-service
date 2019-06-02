import React, {Component, ReactElement} from "react";
import cx from "classnames";
import * as PropTypes from "prop-types";
import {Slider} from "antd";
import {SliderValue} from "antd/lib/slider";

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
  hasMuted: boolean;
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
  timelineDimension: ClientRect | null;
  defaultVolume: number;
  mouseOnPlayhead: boolean;
  startDuration: string;

  constructor(props: AudioPlayerProps) {
    super(props);

    this.state = {
      isPlaying: false,
      playheadPosition: 0,
      activeTimelineWidth: 0,
      currentDuration: "00:00",
      totalDuration: this.formatTime(0),
      hasMuted: false
    };

    this.timeline = null;
    this.playhead = null;
    this.defaultVolume = 0.3;
    this.mouseOnPlayhead = false;
    this.player = null;
    this.timelineDimension = null;
    this.startDuration = "00:00";
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
    if (!seconds) return this.startDuration;
    const minutes = Math.floor(seconds / 60);
    const minutesStr = minutes >= 10 ? minutes.toString() : "0" + minutes;
    seconds = Math.floor(seconds % 60);
    const secondsStr = seconds >= 10 ? seconds.toString() : "0" + seconds;
    return minutesStr + ":" + secondsStr;
  };

  getPosition = (el: HTMLElement): ClientRect => {
    return el.getBoundingClientRect();
  };

  handleEnded = (): void => {
    if (!this.player) return;
    this.setState({...this.state, currentDuration: this.startDuration, isPlaying: false});
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
    this.movePlayhead(e);
  };

  handleTimeUpdate = (): void => {
    if (!this.player || !this.timelineDimension) return;

    const {duration, currentTime} = this.player;
    const {width} = this.timelineDimension;
    const playPercent = width * (currentTime / duration);

    this.setState({
      ...this.state,
      playheadPosition: playPercent,
      activeTimelineWidth: playPercent,
      currentDuration: this.formatTime(currentTime),
    });
  };

  handleCanPlayThrough = (): void => {
    if (!this.player) return;
    this.setState({totalDuration: this.formatTime(this.player.duration)})
  };

  handleVolumeToggle = (): void => {
    if (!this.player) return;
    const {hasMuted} = this.state;
    this.player.muted = !hasMuted;
    this.setState({hasMuted: !hasMuted});
  };

  handleVolumeChange = (value: SliderValue): void => {
    if (!this.player) return;
    if (value instanceof Array) return;

    this.player.volume = value;
  };

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
    if (!this.timeline || !this.player) return;

    const {left, width: timelineWidth} = this.getPosition(this.timeline);
    // const { duration } = this.state;
    const newMargLeft = event.clientX - left;
    let playheadPosition = newMargLeft;

    if (newMargLeft < 0) {
      playheadPosition = 0;
    }
    if (newMargLeft > timelineWidth) {
      playheadPosition = timelineWidth;
    }

    const currentTime: number = this.player.duration * (playheadPosition / timelineWidth);
    const currentDuration: string = this.formatTime(currentTime);

    this.player.currentTime = currentTime;
    this.setState({
      playheadPosition,
      activeTimelineWidth: playheadPosition,
      currentDuration,
    });
  };

  render(): ReactElement {
    const {
      isPlaying,
      playheadPosition,
      activeTimelineWidth,
      currentDuration,
      totalDuration,
      hasMuted,
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
                    className={cx(
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
              <div className="ap__volume__icon" onClick={this.handleVolumeToggle}>
                <i className={cx({"icon-volume-2": !hasMuted, "icon-volume-off": hasMuted})}/>
              </div>
              <div className="ap__volume__controls">
                <Slider
                    defaultValue={this.defaultVolume}
                    tooltipVisible={false}
                    min={0}
                    max={1}
                    step={0.1}
                    onChange={this.handleVolumeChange}
                />
              </div>
            </div>
          </div>
        </div>
    );
  }

  setTimelineRef = _ref => {
    this.timeline = _ref;
    if (!this.timeline) return;
    this.timelineDimension = this.getPosition(this.timeline);
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