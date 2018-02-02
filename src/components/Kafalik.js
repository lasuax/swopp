import React, { Component } from "react";
import { Dropdown, Header } from "semantic-ui-react";

export class Kafalik extends Component {
  render() {
    return (
      <div>
        <Header as="h2">Swopp</Header>
        <Dropdown
          text="Kategori"
          icon="filter"
          floating
          labeled
          button
          className="icon"
        >
          <Dropdown.Menu>
            <Dropdown.Header icon="tags" content="Kategoriler" />
            <Dropdown.Item>Important</Dropdown.Item>
            <Dropdown.Item>Announcement</Dropdown.Item>
            <Dropdown.Item>Discussion</Dropdown.Item>
          </Dropdown.Menu>
        </Dropdown>
      </div>
    );
  }
}
