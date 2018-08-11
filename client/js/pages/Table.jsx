import React from 'react'
import ReactTable from 'react-table'
import 'react-table/react-table.css'
import { Div } from 'glamorous'
import Form from './Form'
import globals from '../../js/globals'

class Table extends React.PureComponent {
    constructor(props) {
        super(props);

        this.state = {
            data: [],
            error: null,
            loading: false
        };
    }

    // buildURLQueryString builds a string containing our form values as a URL parameter
    buildURLQueryString = (params) => {
        let esc = encodeURIComponent;
        let query = Object.keys(params)
            .map(k => esc(k) + '=' + esc(params[k]))
            .join('&');

        return query
    }

    // refreshList triggers a data reload from the API
    refreshList = (params) => {
        let query = this.buildURLQueryString(params)
        this.setState({loading: true})
        this.setState({
            filters: query
        }, this.refReactTable.fireFetchData);
    }

    render() {
        const columns = [{
            Header: 'Provider Name',
            accessor: 'name'
        }, {
            Header: 'Provider Street Address',
            accessor: 'street_address'
        }, {
            Header: 'Provider City',
            accessor: 'city'
        }, {
            Header: 'Provider Zip Code',
            accessor: 'zip_code'
        }, {
            Header: 'Provider State',
            accessor: 'state'
        }, {
            Header: 'Hospital Referral Region Description',
            accessor: 'hrrd'
        }, {
            Header: 'Total Discharges',
            accessor: 'total_discharges'
        }, {
            Header: 'Average Covered Charges',
            accessor: 'average_covered_charges'
        }, {
            Header: 'Average Total Payments',
            accessor: 'average_total_payments'
        }, {
            Header: 'Average Medicare Payments',
            accessor: 'average_medicare_payments'
        }]

        return (
            <div>
                <Div className="col-md-2">
                    <Form
                        handleSubmit={this.refreshList}
                        filters={globals.PROVIDER_FIELDS}
                    />
                </Div>
                <Div className="col-md-9 jumbotron" style={{float: 'right'}}>
                    <ReactTable
                        ref={(refReactTable) => {this.refReactTable = refReactTable;}}
                        sortable={true}
                        data={this.state.data}
                        columns={columns}
                        pageSizeOptions={[5, 20, 100, 200, 300, 500]}
                        filterable={false}
                        manual={false}
                        loading={this.state.loading}
                        loadingText={"loading..."}
                        className={'-striped -highlight'}
                        onFetchData={(s) => {
                            this.setState({loading: true})
                            fetch(globals.PROVIDERS_ENDPOINT + this.state.filters)
                                .then(response => response.json())
                                .then((result) => {
                                    if (result) {
                                        this.setState({data: result, loading: false})
                                    } else {
                                        this.setState({data: [], loading: false})
                                    }
                                })
                                .catch(e => this.setState({
                                    error: e,
                                    data: []
                                }));
                        }}
                    />
                    <Div>
                        <a onClick={()=>this.props.showFilter(false)} className="btn btn-primary btn-lg btn-login btn-block">End Search</a>
                    </Div>
            </Div>
            </div>
        );
    }
}

export default Table
