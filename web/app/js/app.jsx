import React from 'react';
import ReactDOM from 'react-dom';
import { Button } from 'react-bootstrap';

require('../styles/common/buttons.styl');

// JSX
const button = (
    <Button bsStyle="success" bsSize="large">Click me!</Button>
);
ReactDOM.render(button, document.getElementById('button'));
