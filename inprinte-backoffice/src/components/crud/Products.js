import {List, Datagrid, TextField, Edit, SimpleForm, TextInput, Create, SimpleFormIterator, EditButton, ArrayInput, DeleteButton} from 'react-admin';
  
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
  
  export const ProductsEdit = props => (
    <Edit {...props}>
      <SimpleForm>
        <TextInput source="name" label="Nom du produit"/>
        <TextInput source="description" label="Description"/>
        <TextInput source="price" label="Prix"/>
        <TextInput source="category" label="Catégorie"/>
        <ArrayInput source="pictureUrl" label="Liens du produit">
          <SimpleFormIterator>
          <TextInput />
          </SimpleFormIterator>
        </ArrayInput>
        <ArrayInput source="product_files" label="Liens du produit">
          <SimpleFormIterator>
          <TextInput />
          </SimpleFormIterator>
        </ArrayInput>
        <TextInput source="pending_validation" label="En attente"/>
        <TextInput source="is_alive" label="Actif"/>
      </SimpleForm>
    </Edit>
  );

  export const ProductsCreate = props => (
    <Create {...props}>
      <SimpleForm>
        <TextInput source="name" label="Nom du produit"/>
        <TextInput source="description" label="Description"/>
        <TextInput source="price" label="Prix"/>
        <TextInput source="category" label="Categorie"/>
        <TextInput source="product_files" label="Liens du produit"/>
        <TextInput source="product_picture" label="Images du produit"/>
      </SimpleForm>
    </Create>
  );
