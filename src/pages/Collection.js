import React, { Component } from 'react';
import CardGrid from '../components/CardGrid';

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
        <CardGrid cards={cards} />
      </div>
    );
  }
}

export default CollectionPage;
