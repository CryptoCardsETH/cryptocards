import PropTypes from 'prop-types';
import React, { Component } from 'react';
import { LineChart, Line, XAxis, ResponsiveContainer, Tooltip } from 'recharts';
import YAxis from 'recharts/lib/cartesian/YAxis';

class TimeGraph extends Component {
  /**
   * Creates a new Array of data, but inserts missing intervals with a value of 0.
   *
   * @param {any} data - Array of json objects containing xAxisKey and yAxisKey.
   * @param {String} xAxisKey - key used to identify the x-Axis or date of the data.
   * @param {String} yAxisKey - key used to identify the y-Axis or value of the data.
   * @returns []
   * @memberof TimeGraph
   */
  addMissingIntervals(data, intervalLength, xAxisKey, yAxisKey) {
    if (!data || data.length < 2) return [];

    // get all intervals between first and last date in data
    let allIntervals = this.getAllIntervals(
      data[0][xAxisKey],
      data[data.length - 1][xAxisKey],
      intervalLength
    );

    let i = 0;
    let j = 0;
    while (i < allIntervals.length) {
      // check if date from data matches the interval date
      let tmpDate = new Date(data[j][xAxisKey]);
      if (
        j < data.length &&
        allIntervals[i].date.getTime() === tmpDate.getTime()
      ) {
        // store data value in interval array
        allIntervals[i][yAxisKey] = data[j][yAxisKey];
        j++;
      } else {
        // interval is missing in data, so set interval value to 0.
        allIntervals[i][yAxisKey] = 0;
      }

      // set better looking date as xAxisKey
      allIntervals[i][xAxisKey] = this.formatDate(allIntervals[i].date);
      i++;
    }
    return allIntervals;
  }

  /**
   * Gets an Array of json objects with every date inbetween the two dates.
   *
   * @param {string} earliestDate - format (yyyy-mm-dd)
   * @param {string} latestDate - format (yyyy-mm-dd)
   * @param {number} intervalLength - length between intervals in milliseconds
   * @returns []
   * @memberof TimeGraph
   */
  getAllIntervals(earliestDate, latestDate, intervalLength) {
    earliestDate = new Date(earliestDate);
    latestDate = new Date(latestDate);
    let range = latestDate - earliestDate;
    let numIntervals = range / intervalLength;
    return Array.from({ length: numIntervals + 1 }, (v, i) => {
      return { date: new Date(earliestDate.valueOf() + intervalLength * i) };
    });
  }

  formatDate(date) {
    return `${date.getFullYear()}-${this.pad(date.getMonth() + 1)}-${this.pad(
      date.getUTCDate()
    )}`;
  }

  pad(number) {
    return ('0' + number).slice(-2);
  }

  render() {
    let data = this.props.data;
    let intervalLength = this.props.intervalLength;
    let yAxisKey = this.props.yAxisKey;
    let xAxisKey = this.props.xAxisKey;
    let yAxisLabel = this.props.yAxisLabel;
    return (
      <div className="time-graph">
        <ResponsiveContainer width="100%" minHeight="500px">
          <LineChart
            height={400}
            data={this.addMissingIntervals(
              data,
              intervalLength,
              xAxisKey,
              yAxisKey
            )}
            margin={{ top: 5, right: 30, left: 20, bottom: 100 }}
          >
            <XAxis dataKey={xAxisKey} tick={{ angle: 90, dy: 40 }} />
            <YAxis
              label={{
                value: yAxisLabel,
                angle: -90,
                position: 'insideLeft'
              }}
            />
            <Tooltip />
            <Line type="monotone" dataKey={yAxisKey} stroke="#8884d8" />
          </LineChart>
        </ResponsiveContainer>
      </div>
    );
  }
}

TimeGraph.propTypes = {
  data: PropTypes.array.isRequired,
  intervalLength: PropTypes.number.isRequired,
  yAxisKey: PropTypes.string.isRequired,
  xAxisKey: PropTypes.string.isRequired,

  yAxisLabel: PropTypes.optionalString
};

export default TimeGraph;
