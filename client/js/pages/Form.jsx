import React from 'react'
import { Div } from 'glamorous'
import globals from '../../js/globals'

class Form extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            formData: []
        };

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChange(event) {
        const target = event.target;
        let formData = this.state.formData
        formData[target.name] = target.value

        this.setState({formData: formData});
    }

    handleSubmit(event) {
        this.props.handleSubmit(this.state.formData)
        event.preventDefault();
    }

    // renderStateInputField generates a select input of US states
    renderStateInputField = (field) => {
        let optionItems = globals.US_STATES_OPTIONS.map((state) =>
            <option value={state.id}>{state.label}</option>
        )

        return (
            <Div class="form-group" css={{paddingBottom: 5}}>
                <label for={field.name}>{field.label}: </label>
                <select
                    name={field.name}
                    onChange={this.handleChange}
                    style={{float:'right', height: '35px'}}>
                    {optionItems}
                </select>
            </Div>
        )
    }

    // renderFilterInputField generates an input for a filter element
    renderFilterInputField = (field) => {
        const pattern = ('decimal' === field.type) ? "[0-9]+([,\.][0-9]+)?" : "[0-9]*"
        return (
            <Div class="form-group" css={{paddingBottom: 5}}>
                <label for={field.name}>{field.label}: </label>
                <input
                    type="text"
                    pattern={pattern}
                    style={{float:'right'}}
                    size="5"
                    name={field.name}
                    value={this.state.formData[field.name]}
                    onChange={this.handleChange}
                />
            </Div>
        )
    }

    // renderFilterGroup generates a group of inputs from an array
    renderFilterGroup = (filters) => {
       return (
            filters.map((field) => {
                if ('state' === field.type) {
                    return this.renderStateInputField(field)
                }
                return this.renderFilterInputField(field)
            })
        )
    }

    render() {
        return (
            <Div style={{width: 300}}>
                <form onSubmit={e => this.handleSubmit(e)}>
                    <h2>Search Patient Data</h2>
                    <p>One search filter is required</p>

                    {this.renderFilterGroup(this.props.filters)}

                    <Div class="form-group" css={{paddingBottom: 5, paddingTop: 5}}>
                        <a onClick={this.handleSubmit} className="btn btn-primary btn-lg btn-login btn-block">Perform Search</a>
                    </Div>
                </form>
            </Div>
        )
    }
}

export default Form
