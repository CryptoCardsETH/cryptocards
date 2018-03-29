import React, { Component } from 'react';
import '../styles/App.css';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import faExternalLinkAlt from '@fortawesome/fontawesome-free-solid/faExternalLinkAlt';
import 'animate.css';
import { withRouter } from 'react-router-dom';
import { Link } from 'react-router-dom';

class FollowingsList extends Component {
  render() {
    let { followings } = this.props;
    return (
      <div className="row">
        {followings.map((followings, index) => {
          let style = {
            animationDelay: (index % 3) / 10 + 's'
          };
          return (
            <div
              key={index}
              className="following col-lg-3 col-md-5 m-2 animated fadeInUp"
              style={style}
            >
              <div className="overlay-container">
                <img
                  className="card-img-top"
                  src={`http://via.placeholder.com/350?text=following+id+${
                    followings.id
                  }`}
                  alt={followings.nickname}
                />
                <div className="overlay overlay-background">
                  <Link to={'/user/' + followings.id} title="More Details">
                    <FontAwesomeIcon
                      icon={faExternalLinkAlt}
                      className="details"
                      size="1x"
                    />
                  </Link>
                </div>
              </div>
              <div className="following-body">
                <h5 className="follwoing-title text-center">
                  {followings.nickname != null
                    ? followings.nickname
                    : 'Id: ' + followings.id}
                </h5>
              </div>
            </div>
          );
        })}
      </div>
    );
  }
}

export default withRouter(FollowingsList);
