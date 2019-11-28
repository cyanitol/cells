/*
 * Copyright 2007-2017 Charles du Jeu - Abstrium SAS <team (at) pyd.io>
 * This file is part of Pydio.
 *
 * Pydio is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */

'use strict';

exports.__esModule = true;

var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _materialUi = require('material-ui');

var _utilMessagesConsumerMixin = require('../util/MessagesConsumerMixin');

var _utilMessagesConsumerMixin2 = _interopRequireDefault(_utilMessagesConsumerMixin);

var _SortColumns = require('./SortColumns');

var _SortColumns2 = _interopRequireDefault(_SortColumns);

var _ListPaginator = require('./ListPaginator');

var _ListPaginator2 = _interopRequireDefault(_ListPaginator);

/**
 * Specific header for Table layout, reading metadata from node and using keys
 */
exports['default'] = _react2['default'].createClass({
    displayName: 'TableListHeader',

    mixins: [_utilMessagesConsumerMixin2['default']],

    propTypes: {
        tableKeys: _react2['default'].PropTypes.object.isRequired,
        loading: _react2['default'].PropTypes.bool,
        reload: _react2['default'].PropTypes.func,
        dm: _react2['default'].PropTypes.instanceOf(PydioDataModel),
        node: _react2['default'].PropTypes.instanceOf(AjxpNode),
        onHeaderClick: _react2['default'].PropTypes.func,
        sortingInfo: _react2['default'].PropTypes.object
    },

    render: function render() {
        var headers = undefined,
            paginator = undefined;
        if (this.props.node.getMetadata().get("paginationData") && this.props.node.getMetadata().get("paginationData").get('total') > 1) {
            paginator = _react2['default'].createElement(_ListPaginator2['default'], { dataModel: this.props.dm, node: this.props.node });
        }
        return _react2['default'].createElement(
            'div',
            { className: 'mui-toolbar toolbarTableHeader' },
            _react2['default'].createElement(_SortColumns2['default'], _extends({ displayMode: 'tableHeader' }, this.props, { columnClicked: this.props.onHeaderClick })),
            _react2['default'].createElement(
                'div',
                { className: "mui-toolbar-group mui-right" },
                paginator,
                _react2['default'].createElement(_materialUi.FontIcon, {
                    key: 1,
                    title: this.context.getMessage('149', ''),
                    className: "mdi mdi-refresh" + (this.props.loading ? " rotating" : ""),
                    onClick: this.props.reload,
                    style: { padding: 16, display: 'block', cursor: 'pointer', fontSize: 24, color: '#9E9E9E' }
                }),
                this.props.additionalActions
            )
        );
    }
});
module.exports = exports['default'];
