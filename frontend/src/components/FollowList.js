import React, { Component } from 'react';
import '../styles/App.css';
import 'animate.css';
import { withRouter } from 'react-router-dom';
import { ListGroup, ListGroupItem } from 'reactstrap';

class FollowList extends Component {
  render() {
    let { follow } = this.props;
    return (
      <div className="column">
        {follow.map((follow, index) => {
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
