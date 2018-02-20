import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import Select from 'react-select';
import 'react-select/dist/react-select.css';
import { setCardFilterText, setCardSortOption } from '../actions/cards';

class Web3Initialization extends React.Component {
  render() {
    let { filterSortKey } = this.props;
    let filter = this.props.card.filters[filterSortKey];
    return (
      <div>
        <input
          value={(filter && filter.text) || ''}
          onChange={e => {
            this.props.setCardFilterText(filterSortKey, e.target.value);
          }}
        />
        <Select
          name="form-field-name"
          value={(filter && filter.sort) || 'id'}
          onChange={e => {
            this.props.setCardSortOption(filterSortKey, e.value);
          }}
          options={[
            { label: 'Card ID', value: 'id' },
            { label: 'Alphabetically', value: 'az' },
            { label: 'Alphabetically (reverse)', value: 'za' }
          ]}
        />
      </div>
    );
  }
}

function mapStateToProps(state) {
  let { card } = state;
  return { card };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators(
    {
      setCardFilterText,
      setCardSortOption
    },
    dispatch
  );
};

export default connect(mapStateToProps, mapDispatchToProps)(Web3Initialization);
