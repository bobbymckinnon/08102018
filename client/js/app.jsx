import React from 'react'
import ReactDOM from 'react-dom'
import Home from './pages/Home'
import Table from './pages/Table'

class App extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            showFilter: false,
        };
    }

    showFilter = (show) => {
        this.setState({showFilter: show})
    }

    render() {
        if (this.state.showFilter) {
            return (<Table showFilter={this.showFilter}/>);
        } else {
            return (<Home showFilter={this.showFilter}/>);
        }
    }
}

ReactDOM.render(<App />, document.getElementById('app'));
