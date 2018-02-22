import React, { Component } from 'react';

export default class FAQ extends Component {
  render() {
    let items = [
      {
        title: 'How do I create an account?',
        body:
          "You'll need to use <a href='https://metamask.io/' target='_blank'/>MetaMask</a>, a chrome/firefox extension that gives you a safe, in-browser wallet. You then can authenticate by signing a message using your wallet's private key. No passwords needed!"
      },
      {
        title: 'How do I deposit ethereum?',
        body:
          "If you open the MetaMask settings, you can see your wallet's address. You can send ETH to that account from any other wallet, or a service or exchange (e.g. CoinBase). We do not ever hold any of your ethereum."
      },
      {
        title: 'How do I withdraw ethereum?',
        body:
          'Just the reverse - you can sent it to a wallet or exchange via MetaMask.'
      },
      {
        title: 'Do I need to provide any personal information?',
        body:
          'No! You can be completely anonymous. You can optionally provide us with an email address for notifications and important updates, as well as a nickname (or a pseudonym!) to identify your profile on global leaderboards.'
      }
    ];
    return (
      <div>
        <h1>FAQ</h1>
        {items.map(i => (
          <div>
            <h4>{i.title}</h4>
            <div dangerouslySetInnerHTML={{ __html: i.body }} />
          </div>
        ))}
      </div>
    );
  }
}
