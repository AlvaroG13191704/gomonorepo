# Monorepo usando Bazel

Este proyecto usa bazel para construir un monorepo con múltiples proyectos.

## Actualizar y configurar

Para actualizar y configurar el proyecto, ejecute el siguiente comando:

```bash
bazel run //:gazelle
```

Para actualizar las dependencias de los repositorios, ejecute el siguiente comando:

```bash
bazel run //:gazelle-update-repos
```

## Construir

Para construir el proyecto, ejecute el siguiente comando:

```bash
bazel build //...
```


## Si desea construir un proyecto específico, ejecute el siguiente comando:

```bash
bazel build //path/to/project:target
```