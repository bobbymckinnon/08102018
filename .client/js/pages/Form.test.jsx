import React from 'react';
import { shallow } from 'enzyme';
import Form from '../../../client/js/pages/Table';
import { ReactForm } from 'react-form'


describe('Form', () => {
    it('should render Form correctly', () => {
        const component = shallow(<Form />);
        expect(component).toMatchSnapshot();
    });

    it("always receives props length 1", () => {
        const component = shallow(<Form />);
        expect(Object.keys(component.props()).length).toBe(1);
    });
});
