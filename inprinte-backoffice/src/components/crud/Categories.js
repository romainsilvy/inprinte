import {List, Datagrid, TextField, Edit, SimpleForm, TextInput, Create, EditButton, DeleteButton} from 'react-admin';
  
  export const CategoriesList = props => (
    <List {...props}>
      <Datagrid rowClick="edit">
        <TextField source="id" label="Id"/>
        <TextField source="name" label="Catégorie"/>
        <TextField source="is_alive" label="Actif"/>
        <EditButton />
        <DeleteButton />
      </Datagrid>
    </List>
  );
  
  export const CategoriesEdit = props => (
    <Edit {...props}>
      <SimpleForm>
        <TextInput source="name" label="Catégorie"/>
        <TextInput source="is_alive" label="Actif"/>
      </SimpleForm>
    </Edit>
  );

  export const CategoriesCreate = props => (
    <Create {...props}>
      <SimpleForm>
        <TextInput source="name" label="Catégorie"/>
        <TextInput source="is_alive" label="Actif"/>
      </SimpleForm>
    </Create>
  );