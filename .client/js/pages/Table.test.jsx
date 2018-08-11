import React from 'react';
import ReactTable from 'react-table';
import { shallow } from 'enzyme';

import Table from '../../../client/js/pages/Table';
import Form from '../../../client/js/pages/Form';

describe('Table', () => {
    it('should render Table correctly', () => {
        const component = shallow(<Table />);
        expect(component).toMatchSnapshot();
    });

    it("always receives props length 1", () => {
        const component = shallow(<Table />);
        expect(Object.keys(component.props()).length).toBe(1);
    });

    it("always renders a `Form`", () => {
        const component = shallow(<Table />);
        expect(component.find(Form).length).toBe(1);
    });

    it("always renders a `ReactTable`", () => {
        const component = shallow(<Table />);
        expect(component.find(ReactTable).length).toBe(1);
    });
});
