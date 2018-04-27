import React, { Component } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchMyNotifications } from '../actions/users';
import {
  Card,
  CardText,
  CardBody,
  Row,
  CardTitle,
  CardSubtitle
} from 'reactstrap';
class Notifications extends Component {
  componentDidMount() {
    this.props.fetchMyNotifications();
  }
  render() {
    // fake user data
    return (
      <div className="container">
        <h1>All Notifications</h1>
        <Row>
          {this.props.user.notifications.map(n => {
            return (
              <Card>
                <CardBody>
                  <CardTitle>{n.type}</CardTitle>
                  <CardSubtitle>{n.created_at}</CardSubtitle>
                  <CardText>
                    <pre>{JSON.stringify(n.data, true, 2)}</pre>
                  </CardText>
                </CardBody>
              </Card>
            );
          })}
        </Row>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return { user: state.user };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators({ fetchMyNotifications }, dispatch);
};
export default connect(mapStateToProps, mapDispatchToProps)(Notifications);
