import React from 'react';
import { shallow } from 'enzyme';

import Home from '../../../client/js/pages/Home';

describe('Home', () => {
    it('should render Home correctly', () => {
        const component = shallow(<Home />);
        expect(component).toMatchSnapshot();
    });

    it("always renders a link to the search page", () => {
        const component = shallow(<Home />);
        const a = component.find("a");
        expect(a.length).toBeGreaterThan(0);
    });
});
