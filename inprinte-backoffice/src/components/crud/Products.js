import {List, Datagrid, SelectInput, BooleanInput, TextField, Edit, SimpleForm, TextInput, Create, SimpleFormIterator, EditButton, ArrayInput, DeleteButton} from 'react-admin';
import React from 'react'

  export const ProductsList = props => (
    <List {...props}>
      <Datagrid rowClick="edit">
        <TextField source="id" label="Id"/>
        <TextField source="name" label="Nom du produit"/>
        <TextField source="description" label="Description"/>
        <TextField source="price" label="Prix"/>
        <TextField source="pending_validation" label="En attente de validation"/>
        <TextField source="is_alive" label="Actif"/>
        <TextField source="category" label="Catégorie"/>
        <TextField source="firstname" label="Prénom"/>
        <TextField source="lastname" label="Nom"/>
        <TextField source="role" label="Role"/>
        <TextField source="id_user" label="Id user"/>
        <TextField source="rate" label="Moyenne des notes"/>
        <EditButton />
        <DeleteButton />
      </Datagrid>
    </List>
  );
  
  export class ProductsEdit extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            items: [],
            Choices: [],
            DataisLoaded: false
        };
    }
    componentDidMount() {
        fetch("http://localhost:8080/api/categoriesFetch")
            .then((res) => res.json())
            .then((json) => {
              const Choices = [];
              json.CategoriesList.map((item) => (
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
      <BooleanInput source="pending_validation" label="En attente de validation"/>
        <TextInput source="name" label="Nom du produit"/>
        <TextInput source="description" label="Description"/>
        <TextInput source="price" label="Prix"/>
        <SelectInput source="category" label="Catégorie" choices={Choices}/>
        <ArrayInput source="pictureUrl" label="Image du produit">
          <SimpleFormIterator>
          <TextInput />
          </SimpleFormIterator>
        </ArrayInput>
        <ArrayInput source="product_files" label="Lien du produit">
          <SimpleFormIterator>
          <TextInput />
          </SimpleFormIterator>
        </ArrayInput>
      </SimpleForm>
    </Edit>
  );      
  }
}

export class ProductsCreate extends React.Component {
  constructor(props) {
      super(props);
      this.state = {
          items: [],
          Choices: [],
          DataisLoaded: false
      };
  }
  componentDidMount() {
      fetch("http://localhost:8080/api/categoriesFetch")
          .then((res) => res.json())
          .then((json) => {
            const Choices = [];
            json.CategoriesList.map((item) => (
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
      <Create  {...this.props}>
      <SimpleForm>
      <BooleanInput source="is_alive" label="Actif"/>
      <BooleanInput source="pending_validation" label="En attente de validation"/>
      <TextInput source="id_user" label="Id user"/>
        <TextInput source="name" label="Nom du produit"/>
        <TextInput source="description" label="Description"/>
        <TextInput source="price" label="Prix"/>
        <SelectInput source="category" label="Catégorie" choices={Choices}/>
        <ArrayInput source="pictureUrl" label="Image du produit">
          <SimpleFormIterator>
          <TextInput />
          </SimpleFormIterator>
        </ArrayInput>
        <ArrayInput source="product_files" label="Lien du produit">
          <SimpleFormIterator>
          <TextInput />
          </SimpleFormIterator>
        </ArrayInput>
      </SimpleForm>
    </Create>
  );
  }
}
