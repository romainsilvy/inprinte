import {List, Datagrid, TextField, EmailField, Edit, SimpleForm, TextInput, Create, EditButton, DeleteButton, BooleanInput, SelectInput } from 'react-admin';
import React from 'react';

  export const UserList = props => (
    <List {...props}>
      <Datagrid rowClick="edit">
        <TextField source="id" label="Id"/>
        <TextField source="firstname" label="Prénom"/>
        <TextField source="lastname" label="Nom"/>
        <TextField source="role" label="Role"/>
        <EmailField source="email" label="Mail"/>
        <TextField source="phone" label="Téléphone"/>
        <TextField source="is_alive" label="Actif"/>
        <TextField source="address.street" label="Rue" />
        <TextField source="address.city" label="Ville" />
        <TextField source="address.state" label="Région" />
        <TextField source="address.country" label="Pays" />
        <TextField source="address.zipCode" label="Code Postal" />
        <EditButton />
        <DeleteButton />
      </Datagrid>
    </List>
  );

  export class UserEdit extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            items: [],
            Choices: [],
            DataisLoaded: false
        };
    }
    componentDidMount() {
        fetch("http://localhost:8080/rolesFetch")
            .then((res) => res.json())
            .then((json) => {
              const Choices = [];
              json.roleList.map((item) => (
                Choices.push({id: item, name: item})
              ));
              this.setState({
                items: json,
                DataisLoaded: true,
                Choices: Choices
              });
            })
    }
    render() {
        const { Choices } = this.state;
        return (
        <Edit  {...this.props}>
        <SimpleForm>
        <BooleanInput source="is_alive" label="Actif"/>
        <TextInput source="firstname" label="Prénom"/>
        <TextInput source="lastname" label="Nom"/>
        <TextInput source="email" label="Mail"/>
        <TextInput source="phone" label="Téléphone"/>
        <SelectInput source="role" label="Role" choices={Choices}/>
        <TextInput source="address.street" label="Rue" />
        <TextInput source="address.city" label="Ville" />
        <TextInput source="address.state" label="Région" />
        <TextInput source="address.country" label="Pays" />
        <TextInput source="address.zipCode" label="Code Postal" />
      </SimpleForm>
    </Edit>
        )
    }
  }

  export const UserCreate = props => (
    <Create {...props}>
      <SimpleForm>
      <BooleanInput source="is_alive" label="Actif"/>
        <TextInput source="firstname" label="Prénom"/>
        <TextInput source="lastname" label="Nom"/>
        <TextInput source="email" label="Mail"/>
        <TextInput source="phone" label="Téléphone"/>
        <TextInput source="password" label="Mot de passe"/>
        <TextInput source="role" label="Role"/>
        <TextInput source="address.street" label="Rue" />
        <TextInput source="address.city" label="Ville" />
        <TextInput source="address.state" label="Région" />
        <TextInput source="address.country" label="France" />
        <TextInput source="address.zipCode" label="Code Postal" />
      </SimpleForm>
    </Create>
  );
