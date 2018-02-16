import React, { Component } from 'react';
import '../styles/App.css';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import faExternalLinkAlt from '@fortawesome/fontawesome-free-solid/faExternalLinkAlt';
import 'animate.css';

class CardGrid extends Component {
  render() {
    return (
      <div className="row">
        {this.props.cards.map((card, index) => {
          let style = {
            animationDelay: (index % 3) / 10 + 's'
          };
          return (
            <div
              key={index}
              className="card col-lg-3 col-md-5 m-2 animated fadeInUp"
              style={style}
            >
              <div className="overlay-container">
                <img
                  className="card-img-top"
                  src={card.imageUrl}
                  alt={card.title}
                />
                <div className="overlay">
                  <a href="/collection" title="More Details">
                    <FontAwesomeIcon
                      icon={faExternalLinkAlt}
                      className="details"
                      size="2x"
                    />
                  </a>
                </div>
              </div>
              <div className="card-body">
                <h4 className="card-title text-center">{card.title}</h4>
              </div>
            </div>
          );
        })}
      </div>
    );
  }
}

export default CardGrid;
