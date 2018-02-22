import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import Select from 'react-select';
import 'react-select/dist/react-select.css';
import { setCardFilterText, setCardSortOption } from '../actions/cards';
import PropTypes from 'prop-types';

export const FILTER_SORT_ID = { label: 'Card ID', value: 'id' };
export const FILTER_SORT_AZ = { label: 'Alphabetically', value: 'az' };
export const FILTER_SORT_ZA = {
  label: 'Alphabetically (reverse)',
  value: 'za'
};
export const FILTER_SORT_PUBLIC = { label: 'Public only', value: 'public' };
export const FILTER_SORT_HIDDEN = { label: 'Hidden only', value: 'hidden' };

export const FILTER_SORT_PRESET_FULL = [
  FILTER_SORT_ID,
  FILTER_SORT_AZ,
  FILTER_SORT_ZA,
  FILTER_SORT_PUBLIC,
  FILTER_SORT_HIDDEN
];

class CardFilterSort extends React.Component {
  render() {
    let { filterSortKey, sortTypes } = this.props;
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
          value={(filter && filter.sort) || FILTER_SORT_ID.value}
          onChange={e => {
            this.props.setCardSortOption(filterSortKey, e.value);
          }}
          options={sortTypes}
        />
      </div>
    );
  }
}

CardFilterSort.propTypes = {
  filterSortKey: PropTypes.string.isRequired,
  sortTypes: PropTypes.array
};

CardFilterSort.defaultProps = {
  sortTypes: [FILTER_SORT_ID, FILTER_SORT_AZ, FILTER_SORT_ZA]
};

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
export default connect(mapStateToProps, mapDispatchToProps)(CardFilterSort);
