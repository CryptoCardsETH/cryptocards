import React from 'react';

import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import faUserPlus from '@fortawesome/fontawesome-free-solid/faUserPlus';
import faCheck from '@fortawesome/fontawesome-free-solid/faCheck';

const UserDetailFollowing = ({ follow, isFollowing }) =>
  !isFollowing ? (
    <button className="btn btn-success" onClick={follow}>
      <FontAwesomeIcon icon={faUserPlus} /> Follow
    </button>
  ) : (
    <div className="btn bg-primary text-white">
      <FontAwesomeIcon icon={faCheck} /> Following
    </div>
  );

export default UserDetailFollowing;
