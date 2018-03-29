import React, { Component } from 'react';
import '../styles/App.css';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import faExternalLinkAlt from '@fortawesome/fontawesome-free-solid/faExternalLinkAlt';
import 'animate.css';
import { withRouter } from 'react-router-dom';
import { Link } from 'react-router-dom';
import { ListGroup, ListGroupItem } from 'reactstrap';

class FollowList extends Component {
  render() {
    let { follow } = this.props;
    return (
      <div className="column">
        {follow.map((follow, index) => {
          let style = {
            animationDelay: (index % 3) / 10 + 's'
          };
          return (
            <ListGroup>
              <ListGroupItem tag="a" href={'/user/' + follow.id} action>
                {follow.nickname != null ? follow.nickname : 'Id: ' + follow.id}
              </ListGroupItem>
            </ListGroup>
          );
        })}
      </div>
    );
  }
}

export default withRouter(FollowList);
