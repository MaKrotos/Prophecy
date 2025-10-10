/**
 * Утилита для повторных попыток выполнения функций
 */
export const retryWithBackoff = async (fn, maxRetries = 3, options = {}) => {
  const {
    baseDelay = 1000,
    maxDelay = 10000,
    onRetry = null
  } = options;

  let lastError;

  for (let attempt = 0; attempt <= maxRetries; attempt++) {
    try {
      return await fn();
    } catch (error) {
      lastError = error;
      
      if (attempt < maxRetries) {
        const delay = Math.min(baseDelay * Math.pow(2, attempt), maxDelay);
        
        if (onRetry) {
          onRetry(attempt + 1, maxRetries + 1, delay, error);
        }
        
        console.log(`⏳ Ожидание ${delay}ms перед следующей попыткой...`);
        await new Promise(resolve => setTimeout(resolve, delay));
      }
    }
  }

  throw lastError;
};

/**
 * Ожидание условия с таймаутом
 */
export const waitForCondition = (condition, timeout = 5000, checkInterval = 50) => {
  return new Promise((resolve, reject) => {
    if (condition()) {
      resolve(true);
      return;
    }

    const startTime = Date.now();
    const check = () => {
      if (condition()) {
        resolve(true);
      } else if (Date.now() - startTime > timeout) {
        reject(new Error(`Condition timeout (${timeout}ms)`));
      } else {
        setTimeout(check, checkInterval);
      }
    };
    
    check();
  });
};