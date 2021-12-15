/// <reference path="./global.d.ts" />
// @ts-check

/**
 * Implement the functions needed to solve the exercise here.
 * Do not forget to export them so they are available for the
 * tests. Here an example of the syntax as reminder:
 *
 * export function yourFunction(...) {
 *   ...
 * }
 */

/**
 * 
 * @param {number} timer 
 * @returns {string}
 */
export const cookingStatus = (timer) => {
  switch (timer) {
    case null:
    case undefined:
      return 'You forgot to set the timer.';
    case 0:
      return 'Lasagna is done.'
    default:
      return 'Not done, please wait.'
  }
}

/**
 * 
 * @param {number[]} layers 
 * @param {number} avgPrepTime 
 * @returns {number}
 */
export const preparationTime = (layers, avgPrepTime) => {
  if (avgPrepTime != null) {
    return avgPrepTime * layers.length;
  }
  return layers.length * 2;
}

/**
 * 
 * @param {number[]} layers 
 * @returns {object}
 */
export const quantities = (layers) => {
  let quantityReq = {
    noodles: 0, sauce: 0,
  }

  layers.forEach(layer => {
    if (layer.toString() === 'noodles') {
      quantityReq.noodles += 50;
    } else if (layer.toString() === 'sauce') {
      quantityReq.sauce += 0.2;
    }
  });

  return quantityReq;
}

/**
 * 
 * @param {number[]} friendsList 
 * @param {number[]} myList 
 */
export const addSecretIngredient = (friendsList, myList) => {
  const last = friendsList[friendsList.length - 1];
  myList.push(last);
}

/**
 * 
 * @param {object} recipe 
 * @param {number} portion 
 * @returns {object}
 */
export const scaleRecipe = (recipe, portion) => {

  let scaledRecipe = {};

  for (const entry in recipe) {
    if (Object.hasOwnProperty.call(recipe, entry)) {
      scaledRecipe[entry] = (recipe[entry] / 2) * portion;
    }
  }
  return scaledRecipe;
}