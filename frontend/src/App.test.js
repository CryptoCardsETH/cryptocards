import React from 'react';
import App from './App';

import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';

configure({ adapter: new Adapter() });

it('renders without crashing', () => {
  shallow(<App />); //need to shallow render because redux connect is wonky
});
