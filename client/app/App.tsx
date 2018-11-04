import * as React from 'react';

type State = { message: string };
export default class App extends React.Component<{}, State> {
  constructor(props: {}) {
    super(props);

    this.state = {
      message: 'Loading...',
    };
  }

  componentDidMount() {
    fetch('/api/home')
      .then(x => x.json())
      .then(json => {
        this.setState({
          message: json.message,
        });
      });
  }

  render() {
    const { message } = this.state;

    return <div>{message}</div>;
  }
}
