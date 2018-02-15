import React, { Component } from 'react';
import '../styles/App.css';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import Web3Login from '../components/Web3Login';
import { BooleanStatus } from '../components/Icons';

class CollectionPage extends Component {
  render() {
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
        <h1>Card Collection</h1>

        <div class="row">
          {cards.map((card, index) => {
            let style = {
              animationDelay: (index % 3) / 10 + 's'
            };
            return (
              <div
                key={index}
                class="card col-lg-3 col-md-5 m-2 animated fadeInUp"
                style={style}
              >
                <div class="overlay-container">
                  <img
                    class="card-img-top"
                    src={card.imageUrl}
                    alt="Card image"
                  />
                  <div class="overlay">
                    <a href="#" title="More Details">
                      <i
                        class="fas fa-external-link-alt details"
                        alt="More Details"
                      />
                    </a>
                  </div>
                </div>
                <div class="card-body">
                  <h4 class="card-title text-center">{card.title}</h4>
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
