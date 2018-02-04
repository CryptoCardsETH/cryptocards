import React from 'react';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import faCheckCircle from '@fortawesome/fontawesome-free-solid/faCheckCircle';
import faTimesCircle from '@fortawesome/fontawesome-free-solid/faTimesCircle';

export const StatusCheck = () => (
  <FontAwesomeIcon icon={faCheckCircle} color="green" size="lg" />
);
export const StatusCross = () => (
  <FontAwesomeIcon icon={faTimesCircle} color="red" size="lg" />
);

export const BooleanStatus = ({ bool }) =>
  bool ? <StatusCheck /> : <StatusCross />;
