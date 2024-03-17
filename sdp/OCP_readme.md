# Specification Pattern Guide

The Specification Pattern offers a robust framework for encapsulating business rules or criteria and evaluating objects against them. It is particularly beneficial in scenarios requiring object filtration based on dynamic criteria, enabling the combination and modification of these criteria at runtime.

## Overview

This pattern involves an `Specification` interface, which includes a method `isSatisfied(product *Product) bool`. The method evaluates if a given `product` meets the specification, returning `true` for a match and `false` otherwise.

## Implementations

We have several concrete implementations of specifications:

- **`ColorSpecification`**: Validates if a product's color aligns with the specified color.
- **`SizeSpecification`**: Checks whether a product's size matches the specified size.
- **`ColorSizeSpecification`**: Assesses if both the color and size of a product meet the specified criteria.

These implementations are straightforward, directly comparing the product's properties against the criteria defined within each specification.

## Composite Specifications

A key feature of the Specification Pattern is its support for composite specifications, such as the `AndSpecification`. This allows for combining multiple specifications to form a single composite specification. The `AndSpecification`, for example, combines two specifications and checks if a product satisfies both.

### How It Works

When `isSatisfied(product *Product)` is invoked on an `AndSpecification` instance, it sequentially calls the `isSatisfied` method on its two constituent specifications. The product is considered to meet the composite specification only if it satisfies both individual specifications.

This mechanism leverages the polymorphic nature of the `Specification` interface, enabling seamless interchangeability and composition of specifications. Each concrete specification encapsulates its own evaluation logic, facilitating the construction of complex, composite specifications from simpler, single-criterion specifications.

## Principles Demonstrated

### Open/Closed Principle

The Specification Pattern exemplifies the Open/Closed Principle. Systems can introduce new specifications or composite specifications without altering existing code, thereby remaining open for extension but closed for modification.

### Single Responsibility Principle

Each specification is tasked with a single responsibility: to determine whether a product meets a particular criterion. This clear delineation of responsibilities enhances modularity and maintainability.

## Summary

The Specification Pattern provides a flexible and extensible approach to defining business rules and evaluating objects against them. By enabling the composition of complex specifications from simpler ones, it facilitates the dynamic combination and alteration of business rules, adhering to key software design principles.