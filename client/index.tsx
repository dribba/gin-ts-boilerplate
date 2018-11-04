import * as ReactDOM from 'react-dom';
import * as React from 'react';
import App from './app/App';
import './styles/main.less';

require('whatwg-fetch');

ReactDOM.render(<App />, document.getElementById('app'));
