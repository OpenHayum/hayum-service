import React, {Component} from "react";
import classnames from "classnames";
import * as PropTypes from "prop-types";

import "./audioPlayer.less";

const controlNames = {
  IS_PLAYING: "isPlaying",
  IS_PAUSED: 'isPaused',
  IS_RESET: 'isReset',
};

class AudioPlayer extends Component<any, any> {
  static propTypes = {
    src: PropTypes.string
  };

  static defaultProps = {
    src: "http://www.panacherock.com/downloads/mp3/01_Afterlife.mp3"
  };

  timeline: any;
  timelineWidth: any;
  playhead: any;
  mouseOnPlayhead: any;
  duration: any;
  player: any;
  handleEnded: any;

  constructor(props) {
    super(props);
    this.state = {
      isPlaying: false,
      playheadPosition: 0,
      activeTimelineWidth: 0,
      currentDuration: "00:00",
      totalDuration: this.formatTime(null)
    };

    this.timeline = null;
    this.timelineWidth = null;
    this.playhead = null;
    this.mouseOnPlayhead = false;
    this.duration = null;
    this.player = null;
  }

  componentDidMount() {
    window.addEventListener("mouseup", this.mouseUp, false);
  }

  componentWillUnmount() {
    window.removeEventListener("mouseup", this.mouseUp, false);
    this.player.removeEventListener('canplaythrough', this.handleCanPlayThrough, false);
    this.player.removeEventListener('timeupdate', this.handleTimeUpdate, false);
    this.player.removeEventListener('ended', this.handleEnded, false);
  }

  clickPercent = (event) => {
    const {left, width} = this.getPosition(this.timeline);
    return (event.clientX - left) / width;
  };

  formatTime = seconds => {
    if (!!seconds === false) return "00:00";

    let minutes: any;
    minutes = Math.floor(seconds / 60);
    minutes = minutes >= 10 ? minutes : "0" + minutes;
    seconds = Math.floor(seconds % 60);
    seconds = seconds >= 10 ? seconds : "0" + seconds;
    return minutes + ":" + seconds;
  };

  getPosition = el => {
    return el.getBoundingClientRect();
  };

  handleControlClick = ({target}) => {
    const {isPlaying} = this.state;

    if (isPlaying) {
      this.player.pause();
    } else {
      this.player.play();
    }

    this.setState({
      [target.name]: !isPlaying
    });
  };

  handleTimelineClick = e => {
    // console.info(e);
    this.movePlayhead(e);
  };

  handleTimeUpdate = () => {
    const timelineWidth = 300;
    const playPercent = timelineWidth * (this.player.currentTime / this.duration);

    this.setState({
      playheadPosition: playPercent,
      currentDuration: this.formatTime(this.player.currentTime),
    });
  }

  handleCanPlayThrough = () => {
    this.duration = this.player.duration;
    this.setState({totalDuration: this.formatTime(this.player.duration)})
  }

  mouseDown = () => {
    this.mouseOnPlayhead = true;
    window.addEventListener("mousemove", this.movePlayhead, true);
  };

  mouseUp = event => {
    if (this.mouseOnPlayhead) {
      this.movePlayhead(event);
      window.removeEventListener("mousemove", this.movePlayhead, true);
    }
    this.mouseOnPlayhead = false;
  };

  movePlayhead = event => {
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

  render() {
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
                    name={controlNames.IS_PLAYING}
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
    this.timelineWidth = this.timeline.style.width;
    console.info(this.timeline.getBoundingClientRect());
    this.timeline.addEventListener("click", this.handleTimelineClick, false);
  };

  setPlayheadRef = _ref => {
    this.playhead = _ref;
    this.playhead.addEventListener("mousedown", this.mouseDown, false);
  };

  setPlayerRef = _ref => {
    this.player = _ref;
    this.player.volume = 0.2;
    this.player.addEventListener('canplaythrough', this.handleCanPlayThrough, false);
    this.player.addEventListener('timeupdate', this.handleTimeUpdate, false);
    this.player.addEventListener('ended', this.handleEnded, false);
  };
}

export default AudioPlayer;