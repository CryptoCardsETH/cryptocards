import React, { Component } from 'react';
import '../styles/App.css';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import faExternalLinkAlt from '@fortawesome/fontawesome-free-solid/faExternalLinkAlt';
import 'animate.css';
import { withRouter } from 'react-router-dom';
import { Link } from 'react-router-dom';

class FollowersList extends Component {
  render() {
    let { followers } = this.props;
    return (
      <div className="row">
        {followers.map((followers, index) => {
          let style = {
            animationDelay: (index % 3) / 10 + 's'
          };
          return (
            <div
              key={index}
              className="follower col-lg-3 col-md-5 m-2 animated fadeInUp"
              style={style}
            >
              <div className="overlay-container">
                <img
                  className="card-img-top"
                  src={`http://via.placeholder.com/350?text=follower+id+${
                    followers.id
                  }`}
                  alt={followers.nickname}
                />
                <div className="overlay overlay-background">
                  <Link to={'/user/' + followers.id} title="More Details">
                    <FontAwesomeIcon
                      icon={faExternalLinkAlt}
                      className="details"
                      size="1x"
                    />
                  </Link>
                </div>
              </div>
              <div className="follower-body">
                <h5 className="follower-title text-center">
                  {followers.nickname != null
                    ? followers.nickname
                    : 'Id: ' + followers.id}
                </h5>
              </div>
            </div>
          );
        })}
      </div>
    );
  }
}

export default withRouter(FollowersList);
