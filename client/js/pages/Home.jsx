import React from 'react'
import { Div } from 'glamorous'

class Home extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <Div className="container">
                <Div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
                    <h1>Patient Data</h1>
                    <p>You have been asked to assist in the creation of an internal API & Search page for a national healthcare provider. This provider has a set of inpatient prospective payment systems providers.</p>
                    <a onClick={()=>this.props.showFilter(true)} className="btn btn-primary btn-lg btn-login btn-block">Search Patient Data</a>
                </Div>
            </Div>
        )
    }
}

export default Home
