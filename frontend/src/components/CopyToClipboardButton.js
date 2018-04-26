import React from 'react';
import PropTypes from 'prop-types';
import { CopyToClipboard } from 'react-copy-to-clipboard';
import { Button } from 'reactstrap';

const CopyToClipboardButton = ({ copyText, buttonText }) => (
  <CopyToClipboard text={copyText}>
    <Button color="primary">{buttonText}</Button>
  </CopyToClipboard>
);

CopyToClipboardButton.propTypes = {
  copyText: PropTypes.string.isRequired,
  buttonText: PropTypes.string.isRequired
};
export default CopyToClipboardButton;
