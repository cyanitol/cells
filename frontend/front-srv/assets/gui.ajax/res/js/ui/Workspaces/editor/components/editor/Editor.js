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

// IMPORT
import Pydio from 'pydio'
//import FullScreen from 'react-fullscreen';
import Draggable from 'react-draggable';
import { connect } from 'react-redux';
import { IconButton, Paper } from 'material-ui';
import Tab from './EditorTab';
import Toolbar from './EditorToolbar';
import Button from './EditorButton';
import makeMinimise from './make-minimise';

const { EditorActions } = Pydio.requireLib('hoc');
const MAX_ITEMS = 4;

const { makeTransitionHOC, withMouseTracker, withSelectionControls, withContainerSize } = Pydio.requireLib('hoc');

const styles = {
    selectionButtonLeft: {
        position: "absolute",
        top: "calc(50% - 18px)",
        left: "40px"
    },
    selectionButtonRight: {
        position: "absolute",
        top: "calc(50% - 18px)",
        right: "40px"
    },
    iconSelectionButton: {
        borderRadius: "50%",
        width: "36px",
        height: "36px",
        lineHeight: "36px",
        backgroundColor: "rgb(0, 0, 0, 0.87)",
        color: "rgb(255, 255,255, 0.87)"
    }
}

// MAIN COMPONENT
@makeTransitionHOC({translateY: 800}, {translateY: 0})
@withMouseTracker()
@withSelectionControls()
@connect(mapStateToProps, EditorActions)
export default class Editor extends React.Component {

    constructor(props) {
        super(props)

        const {tabDelete, tabDeleteAll, editorModify, editorSetActiveTab} = props

        this.state = {
            minimisable: false
        }

        this.minimise = () => editorModify({isPanelActive: false})
        this.setFullScreen = () => editorModify({fullscreen: typeof (document.fullscreenElement || document.mozFullScreenElement || document.webkitFullscreenElement || document.msFullscreenElement) !== 'undefined'})

        this.closeActiveTab = (e) => {
            const {activeTab} = this.props

            editorSetActiveTab(null)
            tabDelete(activeTab.id)
        }

        this.close = (e) => {
            editorModify({open: false})
            tabDeleteAll()
        }

        // By default, open it up
        editorModify({isPanelActive: true})
    }

    componentDidMount() {
        DOMUtils.observeWindowResize(this.setFullScreen);
    }

    componentWillUnmount() {
        DOMUtils.stopObservingWindowResize(this.setFullScreen);
    }

    enterFullScreen() {
        if(this.props.onFullBrowserScreen){
            this.props.onFullBrowserScreen();
            return;
        }

        if (this.container.requestFullscreen) {
            this.container.requestFullscreen();
        } else if (this.container.mozRequestFullScreen) {
            this.container.mozRequestFullScreen();
        } else if (this.container.webkitRequestFullscreen) {
            this.container.webkitRequestFullscreen(Element.ALLOW_KEYBOARD_INPUT);
        }
    }

    componentWillReceiveProps(nextProps) {

        if (this.state.minimisable) return

        const {translated} = nextProps

        if (!translated) return

        this.recalculate()

        this.setState({minimisable: true})
    }

    recalculate() {

        const {editorModify} = this.props

        if (!this.container) return

        editorModify({
            panel: {
                rect: this.container.getBoundingClientRect()
            }
        })
    }

    renderChild() {
        const {activeTab, tabs, editorSetActiveTab} = this.props

        const filteredTabs = tabs.filter(({editorData}) => editorData)

        return filteredTabs.map((tab, index) => {
            let style = {
                display: "flex",
                width: (100 / MAX_ITEMS) + "%",
                height: "40%",
                margin: "10px",
                overflow: "hidden",
                whiteSpace: "nowrap"
            }

            if (filteredTabs.length > MAX_ITEMS) {
                if (index < MAX_ITEMS) {
                    style.flex = 1
                } else {
                    style.flex = 0
                    style.margin = 0
                }
            }

            if (activeTab) {
                if (tab.id === activeTab.id) {
                    style.margin = 0
                    style.flex = 1
                } else {
                    style.flex = 0
                    style.margin = 0
                }
            }

            return <Tab key={`editortab${tab.id}`} id={tab.id} style={{...style}} />
        })
    }

    render() {
        const {style, activeTab, isActive, hideToolbar, hideSelectionControls, prevSelectionDisabled, nextSelectionDisabled, onSelectPrev, onSelectNext} = this.props
        const {minimisable} = this.state

        const title = activeTab ? activeTab.title : ""
        const onClose = activeTab ? this.closeActiveTab : this.close
        const onMinimise = minimisable ? this.minimise : null
        const onMaximise = this.maximise

        let parentStyle = {
            display: "flex",
            flex: 1,
            overflow: "hidden",
            width: "100%",
            height: "100%",
            position: "relative"
        }

        if (!activeTab) {
            parentStyle = {
                ...parentStyle,
                alignItems: "center", // To fix a bug in Safari, we only set it when height not = 100% (aka when there is no active tab)
                justifyContent: "center"
            }
        }

        return (
            <div style={{display: "flex", ...style}}>
                <AnimatedPaper ref={(container) => this.container = ReactDOM.findDOMNode(container)} onMinimise={this.props.onMinimise}  minimised={!isActive} zDepth={5} style={{display: "flex", flexDirection: "column", overflow: "hidden", width: "100%", height: "100%", transformOrigin: style.transformOrigin}}>
                    {!hideToolbar && <Toolbar style={{position: "absolute", top: 0, left: 0, right: 0, flexShrink: 0}} title={title} onClose={onClose} onFullScreen={() => this.enterFullScreen()} onMinimise={onMinimise} />}

                    <div className="body" style={parentStyle}>
                        {this.props.transitionEnded && this.renderChild()}
                    </div>

                    {!hideSelectionControls && onSelectPrev && (
                        <Button
                            iconClassName="mdi mdi-chevron-left"
                            style={styles.selectionButtonLeft}
                            iconStyle={styles.iconSelectionButton}
                            disabled={prevSelectionDisabled}
                            onClick={() => onSelectPrev()}
                        />
                    )}
                    {!hideSelectionControls && onSelectNext && (
                        <Button
                            iconClassName="mdi mdi-chevron-right"
                            style={styles.selectionButtonRight}
                            iconStyle={styles.iconSelectionButton}
                            disabled={nextSelectionDisabled}
                            onClick={() => onSelectNext()}
                        />
                    )}
                </AnimatedPaper>
            </div>
        );
    }
};

// ANIMATIONS
const AnimatedPaper = makeMinimise(Paper)

// REDUX - Then connect the redux store
function mapStateToProps(state, ownProps) {
    const { editor, tabs } = state

    const activeTab = tabs.filter(tab => tab.id === editor.activeTabId)[0]

    return  {
        style: {},
        hideToolbar: !ownProps.isNearTop,
        hideSelectionControls: !ownProps.isNearTop && !ownProps.isNearLeft && ! ownProps.isNearRight,
        ...ownProps,
        activeTab,
        tabs,
        isActive: editor.isPanelActive
    }
}
