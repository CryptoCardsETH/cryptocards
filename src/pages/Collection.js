import React, { Component } from 'react';
import '../styles/App.css';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import faExternalLinkAlt from '@fortawesome/fontawesome-free-solid/faExternalLinkAlt';

class CollectionPage extends Component {
  render() {
    // fake card data
    let cards = [
      {
        title: 'one',
        imageUrl:
          'https://storage.googleapis.com/ck-kitty-image/0x06012c8cf97bead5deae237070f9587f8e7a266d/262933.svg'
      },
      {
        title: 'two',
        imageUrl:
          'https://storage.googleapis.com/ck-kitty-image/0x06012c8cf97bead5deae237070f9587f8e7a266d/262933.svg'
      },
      {
        title: 'three',
        imageUrl:
          'https://storage.googleapis.com/ck-kitty-image/0x06012c8cf97bead5deae237070f9587f8e7a266d/262933.svg'
      },
      {
        title: 'four',
        imageUrl:
          'https://storage.googleapis.com/ck-kitty-image/0x06012c8cf97bead5deae237070f9587f8e7a266d/262933.svg'
      },
      {
        title: 'five',
        imageUrl:
          'https://storage.googleapis.com/ck-kitty-image/0x06012c8cf97bead5deae237070f9587f8e7a266d/262933.svg'
      },
      {
        title: 'six',
        imageUrl:
          'https://storage.googleapis.com/ck-kitty-image/0x06012c8cf97bead5deae237070f9587f8e7a266d/262933.svg'
      }
    ];
    return (
      <div>
        <h1>My Collection</h1>

        <div className="row">
          {cards.map((card, index) => {
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
      </div>
    );
  }
}

export default CollectionPage;
